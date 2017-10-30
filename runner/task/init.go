package task

import "sync"

/**
 * docker buildは１タスクのみアクセス可能としておく
 */
var dockerBuildMutex *sync.Mutex

func init() {
	dockerBuildMutex = new(sync.Mutex)
}
