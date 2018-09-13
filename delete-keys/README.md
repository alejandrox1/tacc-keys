This program can list all the keys for a given user and delete keys given their
id (the key id can be obtained by listing them first).


Compile it.
```
go build -o keys
```

To list all the keys for a given user
```
$ ./keys --list --id docker
[
  {
    "message": null,
    "result": {
      "id": 53,
      "key_value": "ssh-rsa AAAAB3NzaC1yc2QABAAACAQDf9EWMMH59VfLlZyA8PevlArnKWVca50+SY9pczXozIwSoJNkxSjcNT9UnPNKmCIMt/HO8lxCM1bdraf95ZEjbTn0BqTFzr+hVtoldQetGgGHAOaOJEu44ErTDFiIEnBxN92CAZB9Ijxm+YdeThXEpPvl/ZRgxU70elcMweDJAgimzgOAlP8oJ9etheIZ1pUs8DraKhmBD+uC7fH6sVmT/rbUN4P+WAj7tjeT2LUBf0QLEa5rfYvJ/JpeZFZZ57TVpmpjjKSUl4CABsm/uIBlEMUBzhdSdJkSCvwiEbl47Nc9rxFM/IkfDYSVMGLt4O415AWk5MCHkjseIsj0p/zqsxOnyUVRuiDVyaHOBMuOs16fPae7pthemo1AZzMay17sAX1jIe69KPiB64c3YAr1uO0KXor/TwsmrM1H//1xk+kJIaxW7Y5fNnvqdxd745z4i440sPQ==",
      "username": "docker",
      "tags": [
        {
          "name": "keyservice-test"
        }
      ],
      "tenant": "TACC_PROD",
      "created": "2018-09-10T19:53:56.266195Z"
    },
    "status": "success",
    "version": "TEST"
  },
  {
    "message": null,
    "result": {
      "id": 54,
      "key_value": "ssh-rsa AAAAaC1yc2EAAAAAQDAy73GT1PeH/VcczYV4wvWedIZoUYRzwk/onXpyHBCwYPQFiFpZSnQtJujVeTigzaSbqLKSYpMP5mVCiHveOW/7X/EFZOXWAxn/OOKvsb1+hmV1SVSzXsVUXoa6bvZUB+IcIcepm4BhMuzmv7thZLnNv8Vb9h7gB/ESQIAJoHk2YGixrZTeCw28eYDOxGq/S0LTf6S5hPFDYoVM7lNyBpRWUSmqT9uJVEEaxsAtV7KorEJ",
      "username": "docker",
      "tags": [
        {
          "name": "f_test"
        }
      ],
      "tenant": "TACC_PROD",
      "created": "2018-09-12T20:02:58.329245Z"
    },
    "status": "success",
    "version": "TEST"
  }
]
```


To delete a key
```
$ ./keys --delete --id 53
{
  "message": "Deleted",
  "result": null,
  "status": "success",
  "version": "TEST"
}
```
