import json
import shutil

d = {}

for i in range(1, 11, 3):
    mat = []
    # ki, i
    for j in range(i):
        mat.append([])
        for k in range(i):
            mat[j].append(1)
    d[f"k{i},{i}"] = mat
    try:
        shutil.rmtree(f"k{i},{i}")
    except:
        pass

mat = []
for i in range(3):
    mat.append([])
    for j in range(3):
        mat[i].append(0)

mat[0][0] = 1
mat[0][1] = 1
mat[1][2] = 1
mat[2][0] = 1
d["weird_case"] = mat

mat = []
for i in range(9):
    mat.append([])
    for j in range(8):
        mat[i].append(0)
mat[0][0] = 1
mat[1][0] = 1
mat[2][1] = 1
mat[3][0] = 1
mat[4][0] = 1
mat[4][1] = 1
mat[4][2] = 1
mat[5][1] = 1
mat[5][2] = 1
mat[6][2] = 1
mat[6][3] = 1
mat[7][1] = 1
mat[7][4] = 1
mat[7][6] = 1
mat[8][1] = 1
mat[8][5] = 1
mat[8][7] = 1
d["textbook_ex"] = mat

print(d)
with open("testGraphs.json", "w") as fp:
    json.dump(d, fp)