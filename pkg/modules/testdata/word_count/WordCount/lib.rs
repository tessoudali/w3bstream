use serde::Serialize;
use std::collections::HashMap;
use std::mem;
use std::os::raw::c_void;
use std::ptr::copy;

#[no_mangle]
pub extern "C" fn alloc(size: usize) -> *mut c_void {
    let mut buf = Vec::with_capacity(size);
    let ptr = buf.as_mut_ptr();
    mem::forget(buf);
    return ptr as *mut c_void;
}

// #[no_mangle]
// pub extern "C" fn append(ptr: *mut c_void, size: usize) -> i32 {
//     let mut sli = unsafe { String::from_raw_parts(ptr as *mut _, size, size) };
//     sli.push_str("hello from wasm!");

//     unsafe { copy(sli.as_ptr(), ptr as *mut u8, sli.len()) }
//     return sli.len() as i32;
// }

// #[no_mangle]
// pub extern "C" fn word_count(ptr: *mut c_void, size: usize) -> i32 {
//     let str = unsafe { String::from_raw_parts(ptr as *mut _, size, size) };
//     let cnt = str.split_whitespace().count();
//     return cnt as i32;
// }

#[no_mangle]
pub extern "C" fn run(ptr: *mut c_void, size: usize) -> i32 {
    let str = unsafe { String::from_raw_parts(ptr as *mut _, size, size) };

    let mut map = HashMap::new();
    for word in str.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }

    let mut arr: Vec<Item> = Vec::new();

    for (k, v) in &map {
        arr.push(Item {
            word: String::from(*k),
            count: *v,
        });
    }
    let json_str = serde_json::to_string(&arr).unwrap();

    unsafe { copy(json_str.as_ptr(), ptr as *mut u8, json_str.len()) }
    return json_str.len() as i32;
}

#[derive(Serialize)]
pub struct Item {
    word: String,
    count: i32,
}
