package pbtest

import (
	"playground/protobuf/core"
	"playground/protobuf/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCPU(t *testing.T) {
	

	t.Parallel()
	x := &pb.CPU{
		BrandName: "Intel",
		NumCores:  8,
		Timestamp: timestamppb.Now(),
	}

	bytes, err := proto.Marshal(x)
	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)
	js := core.ToJSONString(x)
	assert.NotEmpty(t, js)
	t.Logf(
		"Bytes Length:%v, JSON Length:%v, Compression Rate: %.2f%%",
		len(bytes),
		len(js),
		(1-float64(len(bytes))/float64(len(js)))*100,
	)
	t.Log(js)
}

func TestMemory(t *testing.T) {
	t.Parallel()
	x := &pb.Memory{
		Value: 2,
		Unit:  pb.Memory_GIGABYTE,
	}

	bytes, err := proto.Marshal(x)
	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)
	js := core.ToJSONString(x)
	assert.NotEmpty(t, js)
	t.Logf(
		"Bytes Length:%v, JSON Length:%v, Compression Rate: %.2f%%",
		len(bytes),
		len(js),
		(1-float64(len(bytes))/float64(len(js)))*100,
	)
	t.Log(js)
}
