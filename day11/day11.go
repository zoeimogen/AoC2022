package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
	"log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Monkey struct {
    items       []int64
	operation   byte
	op_arg      uint
	test_arg    uint
	acts        [2]uint
	inspections int
}

var (
    monkeys     []Monkey
	sugar       *zap.SugaredLogger
	monkeyLimit int64
)

func (m *Monkey) inspect(item int64, reduceWorry bool) {
	sugar.Debugf("  Monkey inspects an item with a worry level of %d.", item)
	a := int64(m.op_arg)
	if a == 0 {
		a = item
	}
	switch m.operation {
	case '*':
		if (item * a) < item {
			log.Panic("Overflow")
		}
		item = item * a
		sugar.Debugf("    Worry level is multiplied by %d to %d", a, item)
	case '+':
		item = item + a
		sugar.Debugf("    Worry level increases by %d to %d", a, item)
	}

	if reduceWorry {
		item = item / 3;
		sugar.Debugf("    Monkey gets bored with item. Worry level is divided by 3 to %d.", item)
	}

	// This is straying outside of a coding problem into math: We can subtract the product
	// of all possible divisor tests if the number ever gets higher than said product.
	item = item % monkeyLimit

	if item % int64(m.test_arg) == 0 {
		sugar.Debugf("    Current worry level is divisible by %d.", m.test_arg)
		sugar.Debugf("    Item with worry level %d is thrown to monkey %d.", item, m.acts[0])
		monkeys[m.acts[0]].items = append(monkeys[m.acts[0]].items, item);
	} else {
		sugar.Debugf("    Current worry level is not divisible by %d.", m.test_arg)
		sugar.Debugf("    Item with worry level %d is thrown to monkey %d.", item, m.acts[1])
		monkeys[m.acts[1]].items = append(monkeys[m.acts[1]].items, item)
	}

	m.inspections++
}

func (m *Monkey) turn(reduceWorry bool) {
	for _, item := range m.items {
		m.inspect(item, reduceWorry)
	}

	m.items = []int64{}
}

func loadMonkeys(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%v\n", err);
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	monkeyLimit = 1

	for scan.Scan() {
		switch strings.Split(strings.TrimSpace(scan.Text()), " ")[0] {
        case "Monkey":
			monkeys = append(monkeys, Monkey{});
		case "Starting":
			var items []int64

			items_str := strings.Split(scan.Text()[18:], ",")
			for _, s := range items_str {
				i, _ := strconv.Atoi(strings.TrimSpace(s))
			    items = append(items, int64(i))
			}
			monkeys[len(monkeys)-1].items = items
	    case "Operation:":
			op_str := strings.Split(scan.Text()[19:], " ")

			o := op_str[1][0];
			b, _ := strconv.Atoi(op_str[2])

			monkeys[len(monkeys)-1].operation = o
			monkeys[len(monkeys)-1].op_arg = uint(b)
		case "Test:":
			a, _ := strconv.Atoi(scan.Text()[21:])
			monkeys[len(monkeys)-1].test_arg = uint(a)
			monkeyLimit *= int64(a)
		case "If":
			if_str := strings.Split(scan.Text()[7:], " ")
			a, _ := strconv.Atoi(if_str[4])
			if if_str[0] == "true:" {
				monkeys[len(monkeys)-1].acts[0] = uint(a)
			} else {
                monkeys[len(monkeys)-1].acts[1] = uint(a)
			}
		}
	}
}

func runDay11Part(reduceWorry bool, rounds int) int {
	var inspections []int

    for round := 1; round <= rounds; round++ {
		for m := range monkeys {
			sugar.Debugf("Monkey %d:", m)
			monkeys[m].turn(reduceWorry)
		}
        sugar.Debugf("After round %d, the monkeys are holding items with these worry levels:", round)
		for m := range monkeys {
			sugar.Debugf("Monkey %d: %s", m, strings.Trim(strings.Replace(fmt.Sprint(monkeys[m].items), " ", ", ", -1), "[]"))
		}
		
		sugar.Infof("== After round %d ==", round)
		inspections = []int{}
		for m, monkey := range monkeys {
			inspections = append(inspections, monkey.inspections)
			sugar.Infof("Monkey %d inspected items %d times", m, monkey.inspections)
		}
	}

	sort.Ints(inspections)
	result := inspections[len(inspections) - 2] * inspections[len(inspections) - 1]
	return result
}

func runDay11(inputFile string) (int, int) {
	loadMonkeys(inputFile)
	part1 := runDay11Part(true, 20)

	monkeys = []Monkey{}
	loadMonkeys(inputFile)
	part2 := runDay11Part(false, 10000)

    return part1, part2
}

func main() {
	inputFile := flag.String("input", "day11-input.txt", "Problem input file")
	debugOutput := flag.Bool("debug", false, "Enable similar debug output to problem examples")
	infoOutput := flag.Bool("info", false, "Enable info level debugging, round summaries only")

	flag.Parse()

	logCfg := zap.NewDevelopmentConfig()
	if !*debugOutput && !*infoOutput {
		logCfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	} else if !*debugOutput {
		logCfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	logger, _ := logCfg.Build()
	sugar = logger.Sugar()
	defer sugar.Sync()

    part1, part2 := runDay11(*inputFile)
	fmt.Printf("%d\n%d\n", part1, part2)
}
