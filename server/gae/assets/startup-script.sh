##########################
# Timeboosterビルドマシン起動タスク
##########################
# 設定値
export TIMEBOOSTER_TOOL_VERSION=stable

### 実行タスク
cd $HOME

## メタデータ設定
export TIMEBOOSTER_PROJECT_ID=`gcloud config get-value project 2> /dev/null | tr -d "\n"`
export TIMEBOOSTER_API_KEY=`curl "http://metadata.google.internal/computeMetadata/v1/instance/attributes/TIMEBOOSTER_API_KEY" -H "Metadata-Flavor: Google"`
export TIMEBOOSTER_ENDPOINT=`curl "http://metadata.google.internal/computeMetadata/v1/instance/attributes/TIMEBOOSTER_ENDPOINT" -H "Metadata-Flavor: Google"`
export BUILD_MACHINE_ID=`curl "http://metadata.google.internal/computeMetadata/v1/instance/name" -H "Metadata-Flavor: Google"`
export BUILD_MACHINE_ZONE=`curl "http://metadata.google.internal/computeMetadata/v1/instance/zone" -H "Metadata-Flavor: Google"`

## タスク実行に必要なコマンドを取得
if [ -n "`which docker`" ]; then
  echo "installed tools"
else
  apt-get install -y curl git-core zip
  curl -fsSL https://get.docker.com/ | sh
fi

## タスクランナー取得
## 基本的にStableバージョンを利用する
rm ./tbrun
gsutil cp gs://${TIMEBOOSTER_PROJECT_ID}.appspot.com/tools/${TIMEBOOSTER_TOOL_VERSION}/linux/tbrun ./tbrun

## タスク実行
chmod 755 ./tbrun
./tbrun > log.txt
gsutil cp log.txt "gs://${TIMEBOOSTER_PROJECT_ID}.appspot.com/logs/${BUILD_MACHINE_ID}/tbrun.txt"

## タスクキャッシュ削除
bash -c "rm -rf "/root/private/${BUILD_MACHINE_ID}-*" > /dev/null"

## 自害する
echo Y | gcloud compute instances stop $BUILD_MACHINE_ID --zone $BUILD_MACHINE_ZONE
