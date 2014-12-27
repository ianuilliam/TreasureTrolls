TreasureTrolls
==============

Treasure-Stealing Trolls
John Sloan

This program runs a simulation of trolls stealing treasure from each other. The user is asked 
for a number of trolls, n, and then the simulation generates n trolls, numbered 0 to n-1. 
Each troll has a bag of treasure, t, initially worth $1000000 and starts under a bridge at a 
position, p, equal to its id * 1000000.

The trolls take turns moving according to id in ascending order. On a trolls turn, it moves to
position p + r * t, where p is its current position, r is a random float between -2 and 2,
and t is its current amount of treasure. After moving, a troll steals half of the treasure from
both its new neighbors. If it only has one neighbor, it steals all that neighbors treasure.
If a troll loses all its treasure, it is removed from the game. The trolls continue taking
turns until only one remains.
