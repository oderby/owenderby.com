.PHONY: resume site serve

resume:
	go run resume/main.go
	google-chrome --headless --disable-gpu --print-to-pdf file://`pwd`/resume/resume_web.html
	mv output.pdf static/resume/resume.pdf
	google-chrome --headless --disable-gpu --print-to-pdf file://`pwd`/resume/resume.html
	mv output.pdf resume.pdf
	rm -f resume/resume_web.html resume/resume.html

site: resume
	rm -rf docs/*
	hugo -d docs

serve: resume
	hugo serve -D
