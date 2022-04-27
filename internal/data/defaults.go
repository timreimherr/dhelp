package data

import (
	"log"

	"github.com/timreimherr/jhelp/internal/model"
)

var (
	sections = []string{
		"Git",
		"Kubernetes",
	}
	gitInfos = []model.Info{
		{
			Key:   "Create Branch",
			Value: "git checkout -b <branch-name>",
		},
		{
			Key:   "Pull Remote Branch",
			Value: "git checkout --track origin/daves_branch",
		},
	}
	kubInfos = []model.Info{
		{
			Key:   "Show Stashes",
			Value: "git stash push -m \"my_stash\"",
		},
		{
			Key:   "Create Branch",
			Value: "git checkout -b <branch-name>",
		},
		{
			Key:   "Pull Remote Branch",
			Value: "git checkout --track origin/daves_branch",
		},
		{
			Key:   "Get pods",
			Value: "kubectl get pods",
		},
		{
			Key:   "Get services",
			Value: "kubectl get svc",
		},
		{
			Key:   "Tail Pod Logs",
			Value: "kubectl logs -f {pod_name}",
		},
		{
			Key:   "Port-Forward",
			Value: "kubectl port-forward service/service-name -n staging 8080:8080",
		},
	}
)

func createDefaultRecords() {
	for _, v := range sections {
		err := CreateSection(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	gitSection := GetSectionByName("Git")[0]
	for _, v := range gitInfos {
		_, err := AddInfoToSection(gitSection.Id, v.Key, v.Value)
		if err != nil {
			log.Fatal(err)
		}
	}

	kubSection := GetSectionByName("Kubernetes")[0]
	for _, v := range kubInfos {
		_, err := AddInfoToSection(kubSection.Id, v.Key, v.Value)
		if err != nil {
			log.Fatal(err)
		}
	}
}
