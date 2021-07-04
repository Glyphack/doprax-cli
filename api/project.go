package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Code        string `json:"projekt_code"`
	Status      string `json:"status"`

	// NodeHost is hostname of node running application it can be used for further actions on project
	NodeHost string `json:"node"`
}

type Projects []Project

func (c *Client) GetProjects() (*Projects, error) {
	res := &Projects{}
	err := c.sendRequest(viper.GetString("host"), "GET", "projects/", nil, res)
	if err != nil {
		return nil, fmt.Errorf("Cannot get projects list %s", err.Error())
	}
	return res, nil
}

func (c *Client) GetProjectByTitle(projectTitle string) (*Project, error) {
	projects, err := c.GetProjects()
	if err != nil {
		return nil, err
	}

	for _, project := range *projects {
		if project.Title == projectTitle {
			return &project, nil
		}
	}
	return nil, errors.New("Project not found")
}

type PullProjectResponse struct {
	Success bool `json:"success"`
}

func (c *Client) PullProject(project *Project) (*PullProjectResponse, error) {
	res := &PullProjectResponse{}
	err := c.sendRequest(project.NodeHost, "POST", fmt.Sprintf("projects/%s/github/pull/", project.Code), nil, res)
	if err != nil {
		return nil, fmt.Errorf("Cannot Pull project %s", err.Error())
	}
	return res, nil
}

type DeployProjectInput struct {
	Operation string `json:"operation"`
	// Host address of node which runs the project
}

type DeployProjectResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func (c *Client) DeployProject(project *Project, params *DeployProjectInput) (*DeployProjectResponse, error) {
	jsonPayload, err := json.Marshal(params)

	if err != nil {
		log.Fatal(err)
	}
	res := &DeployProjectResponse{}
	err = c.sendRequest(project.NodeHost, "POST", fmt.Sprintf("projects/%s/deploy/main/", project.Code), bytes.NewReader(jsonPayload), res)
	if err != nil {
		return nil, fmt.Errorf("Cannot deploy project %s", err.Error())
	}
	return res, nil
}
