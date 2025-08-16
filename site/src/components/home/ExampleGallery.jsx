import React from "react";
import { motion } from "framer-motion";
import CodeBlock from "./CodeBlock";

export default function ExampleGallery() {
  const yamlExample = `- key: a                  # press Caps Lock + A to launch Slack
  file: /Applications/Slack.app
- key: 1                  # press Caps Lock + 1 to open Calendar
  file: /System/Applications/Calendar.app`;

  const deployExample = `karayaml init                                  # ensure Karabiner config exists
karayaml map a /Applications/Slack.app       # add Slack launcher
karayaml map 1 /System/Applications/Calendar.app  # add Calendar launcher`;

  return (
    <section id="examples" className="bg-surface py-32">
      <div className="max-w-7xl mx-auto px-6">
        <div className="text-center mb-20">
          <motion.h2 
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="text-4xl md:text-5xl font-bold mb-6"
          >
            <span className="text-text">See it in action</span>
          </motion.h2>
          <motion.p 
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ delay: 0.1 }}
            className="text-xl text-text-muted max-w-3xl mx-auto"
          >
            Real example from KaraYAML's repo – two shortcuts to launch Slack and Calendar. Replace paths as needed for your apps.
          </motion.p>
        </div>

        <div className="grid lg:grid-cols-2 gap-12 max-w-6xl mx-auto">
          <motion.div
            initial={{ opacity: 0, x: -20 }}
            whileInView={{ opacity: 1, x: 0 }}
            viewport={{ once: true }}
          >
            <h3 className="text-2xl font-bold mb-6 flex items-center gap-3">
              <span className="w-8 h-8 rounded-full bg-gradient-to-r from-blue-500 to-blue-600 flex items-center justify-center text-white text-sm font-bold">
                1
              </span>
              Configuration
            </h3>
            <CodeBlock 
              code={yamlExample}
              language="yaml"
              filename="shortcuts.yaml"
            />
          </motion.div>

          <motion.div
            initial={{ opacity: 0, x: 20 }}
            whileInView={{ opacity: 1, x: 0 }}
            viewport={{ once: true }}
            transition={{ delay: 0.1 }}
          >
            <h3 className="text-2xl font-bold mb-6 flex items-center gap-3">
              <span className="w-8 h-8 rounded-full bg-gradient-to-r from-green-500 to-green-600 flex items-center justify-center text-white text-sm font-bold">
                2
              </span>
              Deploy
            </h3>
            <CodeBlock 
              code={deployExample}
              language="bash"
            />
          </motion.div>
        </div>

        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.3 }}
          className="bg-secondary/10 border border-secondary/20 rounded-2xl p-8 max-w-4xl mx-auto mt-12"
        >
          <p className="text-text-muted leading-relaxed">
            <strong className="text-text">Result:</strong> The above commands will update your YAML and Karabiner config. Now pressing <kbd className="bg-code-bg px-2 py-1 rounded font-mono text-sm border border-border">Caps Lock + A</kbd> opens Slack, and <kbd className="bg-code-bg px-2 py-1 rounded font-mono text-sm border border-border">Caps Lock + 1</kbd> opens Calendar.
          </p>
        </motion.div>
      </div>
    </section>
  );
}