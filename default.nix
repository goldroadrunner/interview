{ pkgs ? import <nixpkgs> { } } : pkgs.mkShell {
    shellHook = ''
        ${pkgs.coreutils}/bin/mkdir --parents gopath &&
	    export GOPATH=$(${pkgs.coreutils}/bin/pwd)/gopath &&
	    true
    '' ;
    buildInputs = [
        pkgs.go
    ] ;
}