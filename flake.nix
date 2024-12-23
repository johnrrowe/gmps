{
  description = "gMPS: Multi-Party Sessions built on top of gRPC";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/release-24.11";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {inherit system;};
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = [
        pkgs.go-task
        pkgs.go
        pkgs.gopls
        pkgs.protobuf
        pkgs.protoc-gen-go
        pkgs.protoc-gen-go-grpc
      ];

      shellHook = ''
      '';
    };
  };
}
