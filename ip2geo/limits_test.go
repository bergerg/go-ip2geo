package ip2geo

import (
	"testing"
	"time"
)

func TestLimitWhenNumberOfPerformedActionsIsGreaterThanAllowed(t *testing.T) {
	limiter := createLimiter(1)
	shouldBeFalse := limiter.limit()
	shouldBeTrue := limiter.limit()
	if shouldBeFalse || !shouldBeTrue {
		t.Fatal("expected to fail but succeeded anyway")
	}
}

func TestLimitOneSecondPassesZerosTheCounter(t *testing.T) {
	limiter := createLimiter(1)
	limiter.limit()
	shouldBeTrue := limiter.limit()
	time.Sleep(1 * time.Second)
	shouldBeFalse := limiter.limit()
	if shouldBeFalse || !shouldBeTrue {
		t.Fatal("expected to fail but succeeded anyway")
	}
}
