# Go Deeper Menu

This example demonstrates using a menu with nested sub views. In this case we will go from Artists to Paintings to Colors.

```
Artists -> Paintings -> Colors
```

## Pushing onto the NavStack 

The artists component will push paintings onto the Navstack. While Paintings will push Colors onto the Navstack. 

## Poping off the NavStack

When a color is selected it will be popped off the navstack. But it will also emit a `ColorSelected` message. Which the Paintings component should handle and follow a similar pattern popping off the navstack and then emitting a `Painting Selected` msg. Likewise the Artist component will do the same and the menu will have all it's selections.

<!-- <img src="out.gif" alt="gif"/> -->



