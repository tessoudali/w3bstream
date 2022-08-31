use anyhow::Result;
use std::ptr::copy;
use wasmtime::*;

fn main() -> Result<()> {
    let mut store = Store::<()>::default();
    let module = Module::from_file(store.engine(), "./run.wasm")?;
    let instance = Instance::new(&mut store, &module, &[])?;

    // wasm init
    let memory = instance
        .get_memory(&mut store, "memory")
        .ok_or(anyhow::format_err!("failed to find `memory` export"))?;
    let alloc = instance.get_typed_func::<i32, i32, _>(&mut store, "alloc")?;
    let run = instance.get_typed_func::<(i32, i32), i32, _>(&mut store, "run")?;

    // text
    let text = String::from("hello iotex hello");

    // request memory from wasm
    let alloc_ptr = alloc.call(&mut store, text.len() as i32 + 100)?;
    let alloc_ptr_host = unsafe { memory.data_ptr(&store).offset(alloc_ptr as _) };

    // copy data into mem
    let text_bytes = text.as_bytes();
    unsafe { copy(text_bytes.as_ptr(), alloc_ptr_host, text_bytes.len()) }

    // exec word_count
    let new_size = run.call(&mut store, (alloc_ptr, text.len() as i32))?;
    println!("New str len {}", new_size);

    // print out as String
    let out = unsafe { String::from_raw_parts(alloc_ptr_host, new_size as _, new_size as _) };
    println!("{}", out);

    Ok(())
}
