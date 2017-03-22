+++
title = "{{ .Title }}"
id = "{{ .Id }}"
author = "{{ .Author }}"
date = "{{ .Date }}"
date_update = "{{ .DateUpdate }}"
language = "{{ .Language }}"
summary = "{{ .Summary }}"
image = "{{ .Image }}"
image_caption = "{{ .ImageCaption }}"
topics = [{{ range .Tags }} "{{ . }}", {{ end }}]
tags = []
# "Cyber Security", "Machine Learning", "Blockchain", "Mixed Reality", "Smart Living" 
# "Arbeitswelten", "Gender", "Demographie", "Design", "Handel"
# "Identitäten und Werte", "Mobility", "Demokratie / Mitsprache", "Gesundheit / Nachhaltigkeit", "Informationsverhalten" 
# "Länder und Kulturen"
+++

{{ .Content }}
