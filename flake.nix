{
  description = "Nix flake for offByOne";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        sys_pkgs = { aarch64-darwin = "x86_64-darwin"; x86_64-darwin = "x86_64-darwin"; x86_64-linux = "x86_64-linux"; };
        sys_name = sys_pkgs.${system};
        pkgs = nixpkgs.legacyPackages.${sys_name};
        packageName = "OffByOne";
      in
      {
        devShell = (pkgs.buildFHSUserEnv {
          name = "offByOne-dev-env";
          targetPkgs = pkgs: (with pkgs; [
            bazel_5
            go_1_18
            git
            glibc
            gcc
            jdk11
          ]);
        }).env;
      });
}
