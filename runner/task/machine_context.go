package task

import (
	"sync"
	"time"
)

type MachineContext struct {
	/**
	 * シャットダウンを行う時刻
	 *
	 * 初期値は起動時間 + 10分
	 *
	 * 最後にビルドが完了した時刻から待機時間として5分を加算する
	 */
	ShutdownTime time.Time

	/**
	 * 同期オブジェクト
	 */
	mutex    *sync.Mutex
	runTasks int
}

/**
 * シャットダウンまでの待機時間
 */
const shutdownWaitMinute time.Duration = 10

func NewContext() *MachineContext {
	result := &MachineContext{
		ShutdownTime: time.Now().Add(time.Minute * shutdownWaitMinute),
		mutex:        new(sync.Mutex),
	}

	return result
}

/**
 * 実行タスクを１つ追加する
 */
func (it *MachineContext) TaskAdd() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.runTasks++
}

/**
 * 実行タスクを１つ完了させる
 */
func (it *MachineContext) TaskDone() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.runTasks--

	// タスクが完了したら待機時間を増やす
	it.ShutdownTime = time.Now().Add(time.Minute * shutdownWaitMinute)
}

/**
 * シャットダウン時刻に達した場合はtrue
 */
func (it *MachineContext) IsShutdownTime() bool {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	if it.runTasks > 0 {
		// 実行中タスクがある
		return false
	}

	if time.Now().Unix() < it.ShutdownTime.Unix() {
		// まだシャットダウン時刻に達しない
		return false
	}

	return true
}
