import React from "react";
import { Badge } from "@/components/ui/badge";
import { motion } from "framer-motion";
import CodeBlock from "./CodeBlock";

export default function Quickstart() {
  const installCode = "brew install swarupdonepudi/tap/karayaml";
  
  const initCode = "karayaml init";
  
  const editCode = `karayaml edit                # open ~/.kara.yaml in VS Code
#    – or –
karayaml map a /Applications/Slack.app`;

  const findReloadCode = `karayaml find slack         # search mappings by app name (case-insensitive)
karayaml reload             # reapply ~/.kara.yaml and refresh Karabiner`;

  return (
    <section id="quickstart" className="max-w-7xl mx-auto px-6 py-32">
      <div className="text-center mb-20">
        <motion.h2 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-4xl md:text-5xl font-bold mb-6"
        >
          <span className="text-text">Get started in minutes</span>
        </motion.h2>
        <motion.p 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
          className="text-xl text-text-muted max-w-3xl mx-auto"
        >
          Three simple commands to transform your keyboard workflow
        </motion.p>
      </div>

      <div className="max-w-4xl mx-auto space-y-12">
        <motion.div
          initial={{ opacity: 0, x: -20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
        >
          <div className="flex items-center gap-4 mb-6">
            <div className="w-10 h-10 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center text-white font-bold text-lg">
              1
            </div>
            <h3 className="text-2xl font-bold">Install</h3>
          </div>
          <CodeBlock code={installCode} />
        </motion.div>

        <motion.div
          initial={{ opacity: 0, x: -20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
        >
          <div className="flex items-center gap-4 mb-6">
            <div className="w-10 h-10 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center text-white font-bold text-lg">
              2
            </div>
            <h3 className="text-2xl font-bold">Initialize default config</h3>
          </div>
          <CodeBlock code={initCode} />
        </motion.div>

        <motion.div
          initial={{ opacity: 0, x: -20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.2 }}
        >
          <div className="flex items-center gap-4 mb-6">
            <div className="w-10 h-10 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center text-white font-bold text-lg">
              3
            </div>
            <h3 className="text-2xl font-bold">Define your shortcuts</h3>
          </div>
          <CodeBlock code={editCode} />
        </motion.div>

        <motion.div
          initial={{ opacity: 0, x: -20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.25 }}
        >
          <div className="flex items-center gap-4 mb-6">
            <div className="w-10 h-10 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center text-white font-bold text-lg">
              4
            </div>
            <h3 className="text-2xl font-bold">Search & reload</h3>
          </div>
          <CodeBlock code={findReloadCode} />
        </motion.div>

        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.3 }}
          className="bg-accent/10 border border-accent/20 rounded-2xl p-8"
        >
          <div className="flex items-start gap-4">
            <Badge variant="outline" className="bg-accent text-accent-foreground border-accent">
              Note
            </Badge>
            <div className="space-y-2">
              <p className="font-semibold text-text">
                After editing or adding, KaraYAML will auto-merge and reload Karabiner
              </p>
              <p className="text-text-muted text-sm leading-relaxed">
                Make sure <strong>Karabiner-Elements</strong> is installed and running. KaraYAML uses Caps Lock as the Hyper key (⌃⌥⌘⇧) for all shortcuts by default. No manual Karabiner steps required.
              </p>
            </div>
          </div>
        </motion.div>
      </div>
    </section>
  );
}