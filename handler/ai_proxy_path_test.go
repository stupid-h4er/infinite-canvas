package handler

import "testing"

func TestResolveAIProxyPathRoutesArkSeedanceToContentGeneration(t *testing.T) {
	got := resolveAIProxyPath("https://ark.cn-beijing.volces.com/api/v3", "doubao-seedance-2-0-260128", "/videos")
	if got != "/contents/generations/tasks" {
		t.Fatalf("path = %q, want /contents/generations/tasks", got)
	}
}

func TestResolveAIProxyPathRoutesAgentPlanSeedanceToContentGeneration(t *testing.T) {
	got := resolveAIProxyPath("https://ark.cn-beijing.volces.com/api/plan/v3", "doubao-seedance-2.0", "/videos")
	if got != "/contents/generations/tasks" {
		t.Fatalf("path = %q, want /contents/generations/tasks", got)
	}
}

func TestResolveAIProxyPathKeepsNonSeedanceVideoPath(t *testing.T) {
	got := resolveAIProxyPath("https://ark.cn-beijing.volces.com/api/v3", "grok-imagine-video", "/videos")
	if got != "/videos" {
		t.Fatalf("path = %q, want /videos", got)
	}
}
