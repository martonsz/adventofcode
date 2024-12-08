FROM python:3.13-alpine
WORKDIR /workdir
COPY ./ /workdir/
RUN [[ -f requirements.txt ]] && pip install -r requirements.txt
CMD python aoc.py
