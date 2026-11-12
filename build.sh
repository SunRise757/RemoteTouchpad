project_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"

rm -rf $project_dir/build

source /var/home/"$(whoami)"/emsdk/emsdk_env.sh
#/var/home/"$(whoami)"/emsdk/emsdk install 3.1.50
#/var/home/"$(whoami)"/emsdk/emsdk activate 3.1.50

/var/home/"$(whoami)"/Qt/6.7.3/wasm_singlethread/bin/qt-cmake -S $project_dir/frontend -B $project_dir/build/
cmake --build /$project_dir/build/mobile --parallel

go -C backend build -o ../build/backend .

#run
./build/backend
