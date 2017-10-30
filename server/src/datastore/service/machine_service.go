package service

import (
	"api_v1"
	"data"
	"datastore/model"
	"fmt"
	"github.com/eaglesakura/swagger-go-core/swag"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"utils"
)

type googleApiRequestService interface {
	Header() http.Header
}

/**
 * ビルドマシン管理用サービス
 */
type MachineService struct {
	req            data.Context
	ctx            context.Context
	computeService *compute.Service
}

const defaultMachineZone = "us-central1-b"
const defaultMachineCpuNum = 24
const defaultMachineRamGb = 64.0
const defaultMachineStorageGb = 48.0

func authorize(req data.Context, service googleApiRequestService) {
	token, _ := req.GetApp().GetFirebaseServiceAccount().GetServiceAccountToken(req.GetOptions().Context, "https://www.googleapis.com/auth/cloud-platform")
	service.Header().Add("Authorization", "Bearer "+token.AccessToken)
}

func NewMachineService(req data.Context) *MachineService {
	result := &MachineService{
		req: req,
		ctx: req.GetOptions().Context,
	}

	result.computeService, _ = compute.New(urlfetch.Client(req.GetOptions().Context))
	return result
}

func (it *MachineService) initBuildMachine(result *model.BuildMachine) {

	// 起動スクリプトの初期設定
	{
		script, _ := it.req.GetAssets().LoadFile("assets/startup-script.sh")
		result.StartupScript = string(script)
	}

	// デフォルトゾーン指定
	{
		result.Zone = defaultMachineZone
	}
}

func (it *MachineService) LoadMachineInfo() *model.BuildMachine {
	g := it.req.GetGoon()

	result := &model.BuildMachine{
		Id: it.GetMachineId(),
	}

	g.RunInTransaction(func(g *goon.Goon) error {
		err := g.Get(result)
		if err != nil {
			it.initBuildMachine(result)
			g.Put(result)
		}
		return nil
	}, nil)

	return result
}

/**
 * マシンIDを取得する
 */
func getMachineId(req data.Context) string {
	key := *req.GetOptions().Auth.ApiKey
	return utils.ToMD5(key)
}

/**
 * マシンIDを取得する
 *
 * ビルドマシンはAPIキーに対して1つ割り当てられている
 */
func (it *MachineService) GetMachineId() string {
	return getMachineId(it.req)
}

/**
 * マシンの状態を取得する
 */
func (it *MachineService) GetMachineStatus() api_v1.MachineState {
	info := it.LoadMachineInfo()

	service := compute.NewInstancesService(it.computeService)
	instanceGet := service.Get(utils.GetGcpProjectId(), info.Zone, it.GetMachineId())
	authorize(it.req, instanceGet)

	instance, err := instanceGet.Do()
	if err != nil {
		it.req.LogError("Machine state[%v]", err.Error())
		return api_v1.MachineState_None
	} else {
		it.req.LogError("Machine raw status[%v]", instance.Status)
		if instance.Status == "RUNNING" {
			return api_v1.MachineState_Running
		} else {
			return api_v1.MachineState_Shutdown
		}
	}
}

/**
 * ビルドマシン用のメタデータを生成する
 */
func (it *MachineService) newInstanceMetadata() []*compute.MetadataItems {
	result := []*compute.MetadataItems{}

	info := it.LoadMachineInfo()

	// APIキー
	result = append(result, &compute.MetadataItems{Key: "TIMEBOOSTER_API_KEY", Value: it.req.GetOptions().Auth.ApiKey})
	// エンドポイント
	result = append(result, &compute.MetadataItems{Key: "TIMEBOOSTER_ENDPOINT", Value: swag.String("https://" + utils.GetGcpProjectId() + ".appspot.com")})
	// 起動スクリプト
	result = append(result, &compute.MetadataItems{Key: "startup-script", Value: swag.String(info.StartupScript)})

	return result
}

func (it *MachineService) Create(req *api_v1.MachineRequest) (*model.BuildMachine, error) {

	machineId := it.GetMachineId()
	projectId := utils.GetGcpProjectId()
	zone := swag.StringValue(req.Zone)
	cpuNum := swag.Int32Value(req.Cpu) / 2 * 2
	ramGb := swag.Float32Value(req.Ram)
	storageGb := swag.Float32Value(req.Storage)

	if len(zone) == 0 {
		zone = defaultMachineZone
	}
	if cpuNum <= 0 {
		cpuNum = defaultMachineCpuNum
	}
	if ramGb <= 0 {
		ramGb = defaultMachineRamGb
	}
	if storageGb <= 0 {
		storageGb = defaultMachineStorageGb
	}

	service := compute.NewInstancesService(it.computeService)
	reqInstanceSpec := &compute.Instance{
		Name: machineId,
		Zone: "https://www.googleapis.com/compute/v1/projects/" + projectId + "/zones/" + zone,
		//MachineType: "https://www.googleapis.com/compute/v1/projects/" + projectId + "/zones/" + zone + "/machineTypes/custom-24-65536",
		MachineType: "https://www.googleapis.com/compute/v1/projects/" + projectId + "/zones/" + zone + "/machineTypes/" + fmt.Sprintf("custom-%v-%v", cpuNum, int(ramGb*1024)),
		Disks: []*compute.AttachedDisk{
			&compute.AttachedDisk{
				AutoDelete: true,
				Boot:       true,
				DeviceName: machineId,
				Mode:       "READ_WRITE",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "projects/ubuntu-os-cloud/global/images/ubuntu-1604-xenial-v20170502",
					DiskType:    "https://www.googleapis.com/compute/v1/projects/" + projectId + "/zones/" + zone + "/diskTypes/pd-ssd",
					DiskSizeGb:  int64(storageGb),
				},
			},
		},
		CanIpForward: false,
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				Network: "https://www.googleapis.com/compute/v1/projects/" + projectId + "/global/networks/default",
				AccessConfigs: []*compute.AccessConfig{
					&compute.AccessConfig{
						Name: "External NAT",
						Type: "ONE_TO_ONE_NAT",
					},
				},
			},
		},
		Tags: &compute.Tags{
			Items: []string{
				"build-machine",
			},
		},
		Metadata: &compute.Metadata{
			Items: it.newInstanceMetadata(),
		},
		ServiceAccounts: []*compute.ServiceAccount{
			&compute.ServiceAccount{
				Email: utils.GetBuildMachineServiceAccount(),
				Scopes: []string{
					"https://www.googleapis.com/auth/cloud-platform",
				},
			},
		},
		Scheduling: &compute.Scheduling{
			AutomaticRestart:  false,
			OnHostMaintenance: "TERMINATE",
			Preemptible:       true,
		},
	}

	insertCmd := service.Insert(projectId, zone, reqInstanceSpec)
	authorize(it.req, insertCmd)
	if _, err := insertCmd.Do(); err != nil {
		return nil, err
	} else {
		info := it.LoadMachineInfo()
		it.req.GetGoon().RunInTransaction(func(g *goon.Goon) error {
			info.Zone = zone
			_, err := g.Put(info)
			return err
		}, nil)
		return info, nil
	}
}

/**
 * マシンを削除する
 * これは不可逆であるが、ゾーンを変更する等の破壊的変更を行えるようにする
 */
func (it *MachineService) Delete() error {
	info := it.LoadMachineInfo()
	service := compute.NewInstancesService(it.computeService)

	// メタデータを更新する
	buildMachineId := it.GetMachineId()
	projectId := utils.GetGcpProjectId()

	deleteCall := service.Delete(projectId, info.Zone, buildMachineId)
	authorize(it.req, deleteCall)
	if _, err := deleteCall.Do(); err != nil {
		return err
	}

	return nil
}

/**
 * マシンを起動する
 */
func (it *MachineService) Start() error {
	info := it.LoadMachineInfo()
	service := compute.NewInstancesService(it.computeService)

	// メタデータを更新する
	buildMachineId := it.GetMachineId()
	projectId := utils.GetGcpProjectId()

	// メタデータを編集する
	{
		getCmd := service.Get(projectId, info.Zone, buildMachineId)
		authorize(it.req, getCmd)

		instance, err := getCmd.Do()
		if err == nil && instance != nil {
			it.req.LogInfo("Update build-machine[%v] metadata", it.GetMachineId())
			metadata := instance.Metadata
			metadata.Items = it.newInstanceMetadata()

			metadataCmd := service.SetMetadata(projectId, info.Zone, buildMachineId, metadata)
			authorize(it.req, metadataCmd)
			_, err = metadataCmd.Do()
			if err != nil {
				it.req.LogError("Metadata update failed[%v]", err.Error())
				return err
			}
			it.req.LogInfo("Metadata update done[%v]", buildMachineId)
		}

	}

	// マシンを起動する
	{
		startCmd := service.Start(projectId, info.Zone, buildMachineId)
		authorize(it.req, startCmd)
		instance, err := startCmd.Do()
		if instance != nil {
			it.req.LogError("Machine raw status[%v]", instance.Status)
		}
		return err
	}
}

/**
 * マシンを停止する
 */
func (it *MachineService) Stop() error {
	info := it.LoadMachineInfo()

	service := compute.NewInstancesService(it.computeService)
	instanceRequest := service.Stop(utils.GetGcpProjectId(), info.Zone, it.GetMachineId())
	authorize(it.req, instanceRequest)

	instance, err := instanceRequest.Do()
	if instance != nil {
		it.req.LogError("Machine raw status[%v]", instance.Status)
	}
	return err
}

/**
 * 起動スクリプトを更新する
 */
func (it *MachineService) SetStartupscript(script string) error {
	machine := it.LoadMachineInfo()
	machine.StartupScript = script

	g := it.req.GetGoon()

	return g.RunInTransaction(func(g *goon.Goon) error {
		_, err := g.Put(machine)
		return err
	}, nil)
}
