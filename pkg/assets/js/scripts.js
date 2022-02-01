const w_sizes_mapping = ["100%", "1024px", "1536px", "2048px", "3072px", "4096px"];
const h_sizes_mapping = ["1024px", "1536px", "2048px"];

$global_shared_data = {
    "w_size": 0,
    "h_size": 0
}

function size_handler() {
    return {
        w_size: $global_shared_data["w_size"],
        "w_mapped_size": "",
        h_size: $global_shared_data["h_size"],
        "h_mapped_size": "",
        init_size_watcher() {
            console.log("initsizewatcher")
            this.$watch('w_size', () => {
                this.update_display_sizes()
            })
            this.$watch('h_size', () => {
                this.update_display_sizes()
            })
            this.update_display_sizes()
        },
        update_display_sizes() {
            this.w_mapped_size = w_sizes_mapping[this.w_size]
            this.h_mapped_size = h_sizes_mapping[this.h_size]
        },
        update_size() {
            $global_shared_data["w_size"] = this.w_size
            $global_shared_data["h_size"] = this.h_size
            htmx.trigger(htmx.find("#getdeps"), "click")
        }
    }
}

function global_shared_data(key) {
    return $global_shared_data[key]
}