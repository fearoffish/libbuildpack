package v7

import (
	"io/ioutil"
	"path/filepath"
	"time"

	cutlass7 "github.com/cloudfoundry/libbuildpack/cutlass/v7"
	. "github.com/onsi/gomega"
)

func PushApp(app *cutlass7.App) {
	Expect(app.Push()).To(Succeed(), "Failed to push %s", app.Name)
	Eventually(app.InstanceStates, 20*time.Second).Should(Equal([]string{"RUNNING"}))
}

func DestroyApp(app *cutlass7.App) {
	if app != nil {
		app.Destroy()
	}
}

func AddDotProfileScriptToApp(dir string) {
	profilePath := filepath.Join(dir, ".profile")
	Expect(ioutil.WriteFile(profilePath, []byte(`#!/usr/bin/env bash
echo PROFILE_SCRIPT_IS_PRESENT_AND_RAN
`), 0755)).To(Succeed())
}
