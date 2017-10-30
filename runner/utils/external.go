package utils

import (
	"bufio"
	"io"
	"os/exec"
)

type ExternalCommand struct {
	/**
	 * カレントディレクトリ
	 */
	Chdir string

	/**
	 * 実行コマンド
	 */
	Commands []string

	/**
	 * 標準出力のハンドラ
	 */
	Stdout func(stdout string)

	/**
	 * 標準エラーのハンドラ
	 */
	Stderr func(stdout string)
}

func (it *ExternalCommand) Run() error {
	cmd := exec.Command(it.Commands[0], it.Commands[1:]...)

	if it.Chdir != "" {
		cmd.Dir = it.Chdir
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	defer stdout.Close()
	defer stderr.Close()

	cmd.Start()

	// 標準出力 / エラーハンドリング
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				return
			} else if it.Stdout != nil {
				it.Stdout(string(line))
			}
		}
	}()

	//  標準エラー
	go func() {
		reader := bufio.NewReader(stderr)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				return
			} else if it.Stderr != nil {
				it.Stderr(string(line))
			}
		}
	}()

	return cmd.Wait()
}
