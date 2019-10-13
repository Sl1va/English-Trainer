a = open("en.txt", 'r')
b = open("rus.txt", 'r')

out = open("out.txt", 'w')

en = a.readlines()
ru = b.readlines()

for i in range(len(en)):
    en[i] = en[i].replace(' ', '_').lstrip().rstrip()
    ru[i] = ru[i].replace(' ', '_').lstrip().rstrip()

    out.write(en[i] + "$" + ru[i] + ";" + "\n")
out.close()