.PHONY: resume site serve build-resume

build-resume:
	go run resume/main.go

# TODO: find a clean way to automate this that produces nice pdf.
resume: build-resume
	google-chrome resume/resume.html
	google-chrome resume/resume_web.html

site:
	cp resume/resume.pdf static/resume/resume.pdf
	rm -rf docs/*
	hugo -d docs

serve:
	cp resume/resume.pdf static/resume/resume.pdf
	hugo serve -D
