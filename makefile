

rever-list:
	cd ./List && go build && ./List

quick-sort:
	cd ./Sort/quick-sort && go build -o s && ./s

radix-sort:
	cd ./Sort/radix-sort && go build -o s && ./s