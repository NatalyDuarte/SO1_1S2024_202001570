#include <linux/module.h> // THIS_MODULE, MODULE_VERSION, ...
#include <linux/init.h>   // module_{init,exit}
#include <linux/proc_fs.h>
#include <linux/sched/signal.h> // for_each_process()
#include <linux/seq_file.h>
#include <linux/fs.h>
#include <linux/sched.h>
#include <linux/mm.h> // get_mm_rss()

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Nataly Guzman");
MODULE_DESCRIPTION("Informacion cpu");
MODULE_VERSION("1.0");

struct task_struct *task;       // sched.h para tareas/procesos
struct task_struct *task_child; // index de tareas secundarias
struct list_head *list;         // lista de cada tareas


static int escribir_a_proc(struct seq_file *file_proc, void *v)
{
    int running = 0;
    int sleeping = 0;
    int zombie = 0;
    int stopped = 0;
    unsigned long rss;
    unsigned long total_ram_pages;
    
    
    total_ram_pages = totalram_pages();
    if (!total_ram_pages) {
        pr_err("No memory available\n");
        return -EINVAL;
    }
    
    #ifndef CONFIG_MMU
        pr_err("No MMU, cannot calculate RSS.\n");
        return -EINVAL;
    #endif
    
    unsigned long total_cpu_time = jiffies_to_msecs(get_jiffies_64());
    unsigned long total_usage = 0;

    for_each_process(task) {
        unsigned long cpu_time = jiffies_to_msecs(task->utime + task->stime);
        unsigned long cpu_percentage = (cpu_time * 100) / total_cpu_time;
        total_usage += cpu_time;
    }
    //---------------------------------------------------------------------------
    seq_printf(file_proc, "{\n\"cpu_total\":%d,\n", total_cpu_time);
    seq_printf(file_proc, "\"cpu_porcentaje\":%d\n", (total_usage * 100) / total_cpu_time);
    seq_printf(file_proc, "}\n");
    return 0;
}

static int abrir_aproc(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_a_proc, NULL);
}

static struct proc_ops archivo_operaciones = {
    .proc_open = abrir_aproc,
    .proc_read = seq_read
};

static int __init modulo_init(void)
{
    proc_create("cpu_so1_1s2024", 0, NULL, &archivo_operaciones);
    printk(KERN_INFO "Insertar Modulo CPU\n");
    return 0;
}

static void __exit modulo_cleanup(void)
{
    remove_proc_entry("cpu_so1_1s2024", NULL);
    printk(KERN_INFO "Remover Modulo CPU\n");
}

module_init(modulo_init);
module_exit(modulo_cleanup);