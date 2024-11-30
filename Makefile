train_repo: 
	@python3 approach/repo/train.py

test_repo: 
	@python3 approach/repo/test.py


test_external: 
	@python3 approach/external/test.py $(FILE)

test:
	python3.11 test.py