package statuspage

import (
	"time"
	"fmt"
)

type ComponentGroup struct {
	CreatedAt   *time.Time  `json:"created_at,omitempty"`
	Description *string     `json:"description,omitempty"`
	ID          *string     `json:"id,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Position    *int        `json:"position,omitempty"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty"`
	Components  []string    `json:"components,omitempty"`
}

type ComponentGroupsResponse []ComponentGroup

func (c *ComponentGroup) String() string {
	var out string
	line := "================="
	out = fmt.Sprintf("\n%s\nCreated: %s\nName: %s\nID: %s\nComponents: %v\n%s\n",
		line,
		*c.CreatedAt,
		*c.Name,
		*c.ID,
		c.Components,
		line,
	)
	return out
}

type ComponentGroupCreateData struct {
    Name        string
    Components  []string
}


// The StatusPage API expects that URL-encoded arrays have [] appended to the value name,
// hence the extra [] for the components slice.
// Note: we can also use go-querystring to achieve this, but it's more verbose.
func (c *ComponentGroupCreateData) String() string {
    return encodeParams(map[string]interface{}{
        "component_group[name]":                  c.Name,
        "component_group[components][]":          c.Components,
    })
}

func (c *Client) GetAllComponentGroups() ([]ComponentGroup, error) {
	return c.doGetComponentGroups("component-groups.json")
}

func (c *Client) GetComponentGroupByID(id string) (*ComponentGroup, error) {
    cgs, err := c.GetAllComponentGroups()
    if err != nil {
        return nil, err
    }
    for _, cg := range cgs {
        if *cg.ID == id {
            return &cg, nil
        }
    }
    return nil, fmt.Errorf("search error: Component group with ID %s not found", id)
}

func (c *Client) GetComponentGroupByName(name string) (*ComponentGroup, error) {
    cgs, err := c.GetAllComponentGroups()
    if err != nil {
        return nil, err
    }
    for _, cg := range cgs {
        if *cg.Name == name {
            return &cg, nil
        }
    }
    return nil, fmt.Errorf("search error: Component group with name %s not found", name)
}

func (c *Client) doGetComponentGroups(path string) ([]ComponentGroup, error) {
	resp := ComponentGroupsResponse{}
	err := c.doGet(path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) doCreateComponentGroup(path string, group *ComponentGroupCreateData) (*ComponentGroup, error) {
    resp := ComponentGroup{}
    err := c.doPost(path, group, resp)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

func (c *Client) CreateComponentGroup(group *ComponentGroupCreateData) (*ComponentGroup, error) {
    return c.doCreateComponentGroup("component-groups.json", group)
}

func (c *Client) doUpdateComponentGroup(group *ComponentGroup, params fmt.Stringer) (*ComponentGroup, error) {
    resp := ComponentGroup{}
    path := fmt.Sprintf("component-groups/%s.json", *group.ID)
    err := c.doPut(path, params, &resp)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

func (c *Client) updateComponentGroup(group *ComponentGroup) (*ComponentGroup, error) {
    var data ComponentGroupCreateData
    data.Name = *group.Name
    data.Components = group.Components

    ucg, err := c.doUpdateComponentGroup(group, &data)
    if err != nil {
        return nil, err
    }
    return ucg, nil
}

// Updates a component group. Expects a complete ComponentGroup struct - consider using
// one of the GetComponentGroup helper functions.
func (c *Client) UpdateComponentGroup(group *ComponentGroup) (*ComponentGroup, error) {
    return c.updateComponentGroup(group)
}

// DeleteComponentGroup deletes a component group.
// As per the API docs, this endpoint only returns 204 on successful deletion -
// it does not return data.
func (c *Client) DeleteComponentGroup(group *ComponentGroup) (error) {
    path := "component-groups/" + *group.ID + ".json"
    err := c.doDelete(path, nil, nil)
    if err != nil {
        return err
    }
    return nil
}