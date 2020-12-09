import json

d = {}

for i in range(1, 10):
    mat = []
    # ki, i
    for j in range(i):
        mat.append([])
        for k in range(i):
            mat[j].append(1)
    d[f"k{i},{i}"] = mat

print(d)
with open("testGraphs.json", "w") as fp:
    json.dump(d, fp)