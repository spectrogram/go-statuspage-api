package statuspage

import (
	"fmt"
	"net/url"
	"time"
)

type Component struct {
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	PageID      *string    `json:"page_id,omitempty"`
	Position    *int       `json:"position,omitempty"`
	Status      *string    `json:"status,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	GroupID     *string    `json:"group_id,omitempty"`
	Showcase    *bool      `json:"showcase,omitempty"`
}

func (c *Component) String() string {
	var out string
	line := "-----------------"
	out = fmt.Sprintf("\n%s\nCreated: %s\nName: %s\nID: %s\nStatus: %s\n%s\n",
		line,
		*c.CreatedAt,
		*c.Name,
		*c.ID,
		*c.Status,
		line,
	)
	return out
}

type ComponentUpdateData struct {
	Data string
}

func (c *ComponentUpdateData) String() string {
	return c.Data
}

type ComponentCreateData struct {
	Name		string
	Description	string
	GroupID		string
	Showcase	bool
}

func (c *ComponentCreateData) String() string {
	return encodeParams(map[string]interface{}{
		"component[name]":                c.Name,
		"component[description]":         c.Description,
		"component[group_id]": 	          c.GroupID,
		"component[showcase]":            c.Showcase,
	})
}

type ComponentsResponse []Component

func (c *Client) doGetComponents(path string) ([]Component, error) {
	resp := ComponentsResponse{}
	err := c.doGet(path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetAllComponents() ([]Component, error) {
	return c.doGetComponents("components.json")
}

func (c *Client) GetComponentByID(id string) (*Component, error) {
	cs, err := c.GetAllComponents()
	if err != nil {
		return nil, err
	}
	for _, cp := range cs {
		if *cp.ID == id {
			return &cp, nil
		}
	}
	return nil, fmt.Errorf("search error: Component with ID %s not found", id)
}

func (c *Client) GetComponentByName(name string) (*Component, error) {
	cs, err := c.GetAllComponents()
	if err != nil {
		return nil, err
	}
	for _, cp := range cs {
		if *cp.Name == name {
			return &cp, nil
		}
	}
	return nil, fmt.Errorf("search error: Component with name %s not found", name)
}

func (c *Client) doUpdateComponent(comp *Component, params fmt.Stringer) (*Component, error) {
	resp := Component{}
	path := fmt.Sprintf("components/%s.json", *comp.ID)
	err := c.doPatch(path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) updateComponent(comp *Component, attr string) (*Component, error) {
	var newVal string
	switch attr {
	case "name":
		newVal = *comp.Name
	case "description":
		newVal = *comp.Description
	case "status":
		newVal = *comp.Status
	}
	params := &ComponentUpdateData{
		Data: fmt.Sprintf("component[%s]=%s", attr, url.QueryEscape(newVal)),
	}
	uc, err := c.doUpdateComponent(comp, params)
	if err != nil {
		return nil, err
	}
	return uc, nil
}

func (c *Client) UpdateComponentName(comp *Component) (*Component, error) {
	return c.updateComponent(comp, "name")
}

func (c *Client) UpdateComponentStatus(comp *Component) (*Component, error) {
	return c.updateComponent(comp, "status")
}

func (c *Client) UpdateComponentDesc(comp *Component) (*Component, error) {
	return c.updateComponent(comp, "description")
}

func (c *Client) CreateComponent(name, description, groupID string, showcase bool) (*Component, error) {
	newComp := &ComponentCreateData{
		Name: name,
		Description: description,
		GroupID: groupID,
		Showcase: showcase,
	}

	resp := Component{}
	err := c.doPost("components.json", newComp, resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteComponent deletes a component.
// As per the API docs, this endpoint only returns 204 on successful deletion -
// it does not return data.
func (c *Client) DeleteComponent(comp *Component) (error) {
	path := "components/" + *comp.ID + ".json"
	err := c.doDelete(path, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
