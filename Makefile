wercker-dev:
	wercker dev --expose-ports

wercker-build:
	wercker --environment=wercker.env build

wercker-deploy:
	wercker --environment=wercker.env deploy
