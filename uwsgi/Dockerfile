FROM python:3
WORKDIR /code
COPY . .
RUN pip install -r requirements.txt \
	&& pip install uwsgi
EXPOSE 8000 8090
CMD ["uwsgi","uwsgi/uwsgi.ini"]
 
