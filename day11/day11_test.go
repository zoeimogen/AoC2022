package main

import (
	"testing"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func TestDay11(t *testing.T) {
	logCfg := zap.NewDevelopmentConfig()
	logCfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	logger, _ := logCfg.Build()
	sugar = logger.Sugar()
	defer sugar.Sync()

	part1, part2 := runDay11("day11-test.txt")

	if part1 != 10605 {
		t.Errorf("Part 1 test returned %d; want 10605", part1)
	}

	if part2 != 2713310158 {
		t.Errorf("Part 2 test returned %d; want 2713310158", part2)
	}
}
