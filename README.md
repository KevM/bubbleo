# BubbleO

BubbleO is a collection of components for the excellent terminal UI tool [bubbletea](https://github.com/charmbracelet/bubbletea). 

## [Navstack](https://github.com/KevM/bubbleo/blob/main/navstack/model.go)

Add support to your bubble tea application to easily transition between component models. The example below uses the [menu component](https://github.com/KevM/bubbleo/blob/main/menu/model.go) to let a user pick a color from a list of artist paintings.

<img src="examples/deeper/demo.gif" alt="Recording of the deeper example demo"/>

```go 
w := window.New(120, 25, 0, 0)
ns := navstack.New(&w)
m := model{
    // your tea Component goes here
}
ns.Push(navstack.NavigationItem{Model: m, Title: "main menu"})

p := tea.NewProgram(ns, tea.WithAltScreen())
p.Run()
```

Push, well, pushes a new navigation item on the nav stack. The title is used for breadcrumbs (more later). The model at  the top of the navstack has it's Update and View funcs used effectively making it the presented component on the stack.

Popping the stack will remove the topmost navigation item from the stack. 

### Navigation 

Navigation is accomplished by your components when they publish messages like `navstack.PushNavigation{}` or `navstack.PopNavigation{}`.

#### Pushing a new component onto the stack

This example is from the included [menu component](https://github.com/KevM/bubbleo/blob/main/menu/model.go) which presents a list of choices. When a menu item is selected by pressing `enter` the choice's model is pushed onto the stack by publishing `navstack.PushNavigation`.

```go 
    case tea.KeyEnter.String():
        choice, ok := m.list.SelectedItem().(choiceItem)
        if ok {
            m.selected = &choice.key
            item := navstack.NavigationItem{Title: choice.title, Model: choice.key.Model}
            cmd := utils.Cmdize(navstack.PushNavigation{Item: item})
            return m, cmd
        }
```

There is no limit to the depth of the navigation stack. And the stack components may be dynamic based on your application and user needs.

> **Note:** [Cmdize](https://github.com/KevM/bubbleo/tree/main/utils/utils.go) simply wraps the given arg in a `tea.Cmd` (func that returns a `tea.Msg`)

#### Popping the stack

To pop a component off the stack you might do the following in your bubbletea `Update` func. 

```go 
	case color.ColorSelected:
		pop := utils.Cmdize(navstack.PopNavigation{})
		cmd := utils.Cmdize(ArtistSelected{Name: m.Artist.Name, Color: msg.RGB})
		return m, tea.Sequence(pop, cmd)
```

The first cmd pops the current component off the stack while the second command is received by the previous component on the stack. This allows you to communicate the actions taken by the user up the nav stack. 

If there are no items on the stack `tea.Quit` command is sent and the applicaiton exits.

> **Important:** In this example we are using `tea.Sequence` rather than the normal `tea.Batch` to ensure the messages are played in the correct order. This ensures the pop is played before the `ArtistSelected` which means the component below on the stack will get the message after the navstack is update. 

## Menu

The menu component wraps the excellent The [bubble/list component](https://github.com/charmbracelet/bubbletea/tree/master/examples/list-default) to let the user select from a menu of choices. Each choice is a model, title, and optional description which upon selection will take the user to this component's view. Menu expects to be used within a navigation stack. See the [simple](https://github.com/KevM/bubbleo/tree/main/examples/simple) and [deeper](https://github.com/KevM/bubbleo/tree/main/examples/deeper) examples for detailed usage.

## Breadcrumb

It is handy to give the user a sense of place by presenting a [breadcrumb view](https://www.smashingmagazine.com/2022/04/breadcrumbs-ux-design/) showing them where they come from.

Checkout the deeper example for usage of breadcrumbs.


