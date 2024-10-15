# novelgo
Novel Go games written in Go

+ Motivation
Normal Go games with small board sizes have been thoroughly studied, and mature software for full board size is also available.
It'd be intersting to see what it's like if we invent some "novel" Go games by tweaking some of its rules,
for example, by making the board cyclic, so it "wraps" around in all directions.
This eliminates the idea of "edges" and "corners" and the board is now a "cyclic group" of equivalence.
This repo is created to play with such "novel" setups.
At the moment it only has a CLI tool for "cyclic" games, and a basic backend server for CRUD operations.
A front end will be created in the future.
