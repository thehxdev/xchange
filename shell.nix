{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    # nativeBuildInputs is usually what you want -- tools you need to run
    nativeBuildInputs = with pkgs.buildPackages; [
        go
        gopls
        #gdb # for debugging
        #rnix-lsp # for nix lsp
    ];
}
