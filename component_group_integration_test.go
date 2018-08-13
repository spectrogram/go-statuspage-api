// +build integration

package statuspage

import (
    "encoding/json"
    "os"
    "sort"
    "testing"

    "github.com/stretchr/testify/assert"
)

var c *Client

type Credentials struct {
    ApiKey  string  `json:"apiKey"`
    PageID  string  `json:"pageID"`
}

func setup() {
    var creds Credentials
    var err error

    // load credentials from file and create a client
    if f, err := os.Open(".client-testing.json"); err == nil {
        json.NewDecoder(f).Decode(&creds)
        f.Close()
    } else {
        panic("Could not find .client-testing.json")
    }

    c, err = NewClient(creds.ApiKey, creds.PageID)
    if err != nil {
        panic("Error creating new client!")
    }
}

func TestClient_CreateComponentGroup(t *testing.T) {
    // Component groups must be create with components in them
    // We will create a dummy Component
    var newComp ComponentCreateData = ComponentCreateData{
        Name: "StatusPage API Client Test Component (Group)",
        Description: "Component for testing creation of Component Groups",
        GroupID: "",
        Showcase: false,
    }

    _, err := c.CreateComponent(&newComp)
    if err != nil {
        t.Errorf("Error creating a new Component for Component Group testing: %s\n", err)
    }

    comp, err := c.GetComponentByName("StatusPage API Client Test Component (Group)")
    if err != nil {
        t.Fatalf("Error getting new Component for Component Group testing: %s\n", err)
    }

    // Now create a new ComponentGroup containing the Component we just created
    newCompGroupList := []string{*comp.ID}
    var newCompGroup ComponentGroupCreateData = ComponentGroupCreateData{
        Name: "StatusPage API Client Test Component Group",
        Components: newCompGroupList,
    }
    _, err = c.CreateComponentGroup(&newCompGroup)
    if err != nil {
        t.Errorf("Error creating a new Component Group: %s\n", err)
    }

    compGroup, err := c.GetComponentGroupByName("StatusPage API Client Test Component Group")
    if err != nil {
        t.Fatalf("Error getting new Component Group: %s\n", err)
    }

    // Check that the component group contains what we expect
    assert.Equal(t, *compGroup.Name, "StatusPage API Client Test Component Group")
    assert.Equal(t, len(compGroup.Components), len(newCompGroupList))

    // Check that the component group contains exactly the same components
    sort.Strings(compGroup.Components)
    sort.Strings(newCompGroupList)
    for i := 0; i < len(compGroup.Components); i++ {
        if compGroup.Components[i] != newCompGroupList[i] {
            t.Errorf("Component group did not contain the expected components!")
        }
    }
}

func TestClient_DeleteComponentGroup(t *testing.T) {
    // We'll delete the component group we created earlier
    compGroup, err := c.GetComponentGroupByName("StatusPage API Client Test Component Group")
    if err != nil {
        t.Fatalf("Error getting existing Component Group: %s\n", err)
    }

    err = c.DeleteComponentGroup(compGroup)
    if err != nil {
        t.Errorf("Error deleting existing Component Group: %s\n", err)
    }

    // Also delete the component we made for the test
    comp, err := c.GetComponentByName("StatusPage API Client Test Component (Group)")
    if err != nil {
        t.Errorf("Error getting existing Component (from Component Group tests): %s\n", err)
    }

    err = c.DeleteComponent(comp)
    if err != nil {
        t.Fatalf("Error deleting existing Component (from Component Group tests): %s\n", err)
    }

}

func TestMain(m *testing.M) {
    setup()
    r := m.Run()
    os.Exit(r)
}