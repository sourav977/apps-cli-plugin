//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 VMware, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by diegen. DO NOT EDIT.

package v1alpha1

import (
	testingx "testing"

	testing "dies.dev/testing"
)

func TestClusterSupplyChainDie_MissingMethods(t *testingx.T) {
	die := ClusterSupplyChainBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ClusterSupplyChainDie: %s", diff.List())
	}
}

func TestSupplyChainSpecDie_MissingMethods(t *testingx.T) {
	die := SupplyChainSpecBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for SupplyChainSpecDie: %s", diff.List())
	}
}

func TestSupplyChainStatusDie_MissingMethods(t *testingx.T) {
	die := SupplyChainStatusBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for SupplyChainStatusDie: %s", diff.List())
	}
}

func TestWorkloadDie_MissingMethods(t *testingx.T) {
	die := WorkloadBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for WorkloadDie: %s", diff.List())
	}
}

func TestWorkloadSpecDie_MissingMethods(t *testingx.T) {
	die := WorkloadSpecBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for WorkloadSpecDie: %s", diff.List())
	}
}

func TestWorkloadStatusDie_MissingMethods(t *testingx.T) {
	die := WorkloadStatusBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for WorkloadStatusDie: %s", diff.List())
	}
}