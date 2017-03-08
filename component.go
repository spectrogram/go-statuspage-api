package statuspage

import (
	"fmt"
	"time"

	"github.com/kr/pretty"
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
}

type ComponentResponse []Component

func (c *Client) doGetComponents(path string) ([]Component, error) {
	resp := ComponentResponse{}
	err := c.doGet(path, nil, &resp)
	if err != nil {
		return nil, err
	}
	pretty.Println(resp)
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
