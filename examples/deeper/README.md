# A "Deeper" Hierarchical Menu

This example demonstrates using the shell component which wraps the navstack and breadcrumbs in a handy application container. There are 3 levels of menus navigating the user through the application. In this case we will go from _Artists_ to their _Paintings_ to the prominent _Colors_ used in that painting.

```
Artists -> Paintings -> Colors
```

The user can select and preview the desired color and is told which artist and painting it comes from. While this applicaton is simple it demonstrates how each level of navigation could be a more robust complex user experience. The breadcrumbs give the user a sense of place so they do not get lost within a deep hierarchy of actions.

## Pushing onto the NavStack 

The artists component will push paintings onto the navstack. Each Painting will push its Colors onto the navstack. 

## Popping off the NavStack

When a color is selected it will be popped off the navstack. But it will also emit a `ColorSelected` message. Which the Paintings component will handle and follow a similar pattern popping off the navstack and then emitting a `Painting Selected` msg. Likewise the Artist component will do the same and the menu will have all it's selections.

<img src="demo.gif" alt="vhs recording of this TUI example"/>



