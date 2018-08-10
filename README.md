# go-statuspage-api
Go library for StatusPage.io API

## Usage

### Base

```golang
package main

import "github.com/yfronto/go-statuspage-api"

const (
  apiKey = "....."
  pageID = "....."
)

func main() {
  c, err := statuspage.NewClient(apiKey, pageID)
  if err != nil {
    // ...
   }
   // Do stuff
}
```

### Components
```golang
// Pretty print all components on a StatusPage
resp, _ := c.GetAllComponents()
for _, component := range resp {
  fmt.Printf("%v", &component)
}
```

### Component groups
```golang
// Pretty print all component groups on a StatusPage
resp, _ := c.GetAllComponentGroups()
for _, componentGroup := range resp {
  fmt.Printf("%v", &componentGroup)
}
```

```golang
// Update a component group
newName := "foobar"
newComponents := []string{"gibberish", "gobbledegook"}

group, _ := c.GetComponentGroupByID("baz")
group.Components = newComponents
group.Name = newName
resp, _ := c.UpdateComponentGroup(group)
```

```golang
// Delete a component group
func DeleteComponentGroup() {
  group, _ := c.GetComponentGroupByID("deadbeef")
  err = c.DeleteComponentGroup(group)
}
```

Things are still in progress.
