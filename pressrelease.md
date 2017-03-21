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
release_type = "pressrelease"
tags = [{{ range .Tags }} "{{ . }}", {{ end }}]
+++

{{ .Content }}
