# novelgo
Novel Go games written in Go

+ Motivation
Normal Go games with small board sizes have been thoroughly studied,
and mature softwares for the full 19x19 board size are also available.
It's intersting to see what it'd be like if we invent some "novel" Go games by tweaking some of its rules.
For example, by making the board cyclic, so it "wraps" around in all directions.
This eliminates the notion of "edges" and "corners", and the board is now a "cyclic group" of equivalence.
This repo is created to play with such "novel" setups.

+ Status
At the moment it only has a CLI tool for "cyclic" games (full logic with toggle not implemented yet),
which is able to save and read gamplay files.
after adding a basic backend server for CRUD operations, a suitable frontend will also be created in the future.
