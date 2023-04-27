generate:
	kubectl apply -f $(basic).yaml

ingress:
	make basic=ingress generate

dashboard:
	make basic=dashboard generate
	make basic=auth generate

all:
	make dashboard
	make -C minio make
	make -C mysql make
	make -C redis make
