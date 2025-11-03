#!/usr/bin/env python3
with open('Home_minimal_template.txt', 'r') as f:
    template = f.read()

with open('Home_backup_v2.vue', 'r') as f:
    lines = f.readlines()
    script_and_style = ''.join(lines[1337:])  # From line 1338 (index 1337) to end

with open('Home.vue', 'w') as f:
    f.write(template)
    f.write('\n\n')
    f.write(script_and_style)

print("Done! Created Home.vue")
