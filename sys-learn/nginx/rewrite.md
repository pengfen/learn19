```
location / {
	if (!-e $request_filename) {
		rewrite ^/([0-9]+)$ /index.php?m=Show&a=index&roomnum=$1 last;
		rewrite ^(.*)$ /index.php?s=$1 last;
		break;
	}
}
```