# Reinforcement-Learning-Go
Ejercicios resueltos de "Reinforcement Learning: An Introduction de Richard S. Sutton and Andrew G. Barto", programados en el lenguaje Go

# Gambler's Problem
Resultados.<br />
con ph = 0.4:<br /> 
![valuefunctionfound0 4](https://user-images.githubusercontent.com/6053293/38634397-3ada7b44-3d77-11e8-91a0-0b4412a396b3.png)<br />
![finalpolicy0 4](https://user-images.githubusercontent.com/6053293/38634432-534fde9e-3d77-11e8-9bda-455114862e25.png)<br />

con ph = 0.25:<br />
![valuefunctionfound0 25](https://user-images.githubusercontent.com/6053293/38634474-6e839778-3d77-11e8-9f71-a33ee0a080c7.png)<br />
![finalpolicy0 25](https://user-images.githubusercontent.com/6053293/38634472-6dc8f878-3d77-11e8-9522-cb7073d610a6.png)<br />

con ph = 0.55:<br />
![valuefunctionfound0 55](https://user-images.githubusercontent.com/6053293/38634536-9b75ed9e-3d77-11e8-9d34-184242c87dcd.png)<br />
![finalpolicy0 55](https://user-images.githubusercontent.com/6053293/38634535-9b5a82c0-3d77-11e8-9fe2-5f096ab16e3d.png)<br /><br />

# Windy GridWorld
Solución al ejercicio 6.5 <br />
D | D | U | R | R | R | R | R | R | D |
U | R | R | R | R | R | R | R | R | D |
R | R | R | L | R | R | R | L | R | D |
U | U | R | R | R | R | R | G | R | D |
D | R | R | U | R | R | D | D | L | L |
R | R | R | R | R | R | L | D | L | L |
R | R | R | U | R | D | R | L | U | L |
0   0   0   1   1   1   2   2   1   0   <- wind <br /><br />

Solución al ejercicio 6.9 Windy Gridworld with King’s Moves <br />
R | L | R | R | DR| R | UR| UR| UR| D |
D | DR| DR| R | UR| UR| DR| R | R | D |
DR| D | D | R | R | R | DR| DR| R | DL|
DL| D | DL| DR| R | U | R | G | DL| D |
DL| DR| DR| DL| DL| DR| DR| D | L | L |
DR| UR| U | DR| DR| DR| R | D | UL| L |
DR| R | DR| DR| UR| UL| R | UL| D | UL|
0   0   0   1   1   1   2   2   1   0   <- wind <br /><br />

Solución al ejercicio 6.10 Stochastic Wind <br />
UR| UR| R | DR| UR| DR| UR| U | R | DR|
UR| UR| UR| DR| UR| R | DR| R | D | DR|
UR| UR| DR| UR| UR| UR| R | D | DL| L |
DR| R | D | R | UR| R | DR| G | L | DL|
R | DR| DR| UR| R | UR| R | DR| UL| L |
DR| DR| R | DR| DR| DR| DR| D | DL| UL|
R | R | UR| UR| D | DR| DL| D | L | U |
0   0   0   1   1   1   3   3   0   0   <- wind <br /><br />


