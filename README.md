# Treasure Hunt

```
########
#......#
#.###..#
#...#.##
#X#...E#
########
```

Legends :
- `#`: Obstacle
- `.`: Clear Path
- `X`: Initial Player starting positing
- `E`: Improvement, target location
- `$`: Possible treasure location

## Tracking Direction to the Location Target
As an example, given **E** is a target location. To get direction until **E**, we use BFS (Breadth First Search) Algorithm.

## Get All Possibility Treasure Located
To get all possible treasure hunt located, we can just check the grid cell that have value `#`, `X` or `E`.

