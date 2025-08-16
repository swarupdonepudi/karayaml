
import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Terminal, Flag } from "lucide-react";
import { motion } from "framer-motion";
import CodeBlock from "./CodeBlock";

const commands = [
  {
    command: "karayaml init",
    description: "one-time setup (creates default Karabiner config if needed)"
  },
  {
    command: "karayaml edit",
    description: "open the shortcuts YAML in your editor (validates on save)"
  },
  {
    command: "karayaml map <key> <path_to_app>",
    description: "map a key to an app (shortcut mapping)"
  },
  {
    command: "karayaml list",
    description: "list all configured shortcuts in a table"
  },
  {
    command: "karayaml version",
    description: "show KaraYAML CLI version"
  }
];

const flags = [
];

export default function CLIReference() {
  const codeExample = `karayaml init           # one-time setup (creates default Karabiner config if needed)
karayaml edit           # open the shortcuts YAML in your editor (validates on save)
karayaml map <key> <path_to_app>   # map a key to an app
karayaml list           # list all configured shortcuts in a table
karayaml version        # show KaraYAML CLI version`;

  return (
    <section id="cli" className="max-w-7xl mx-auto px-6 py-32">
      <div className="text-center mb-20">
        <motion.h2 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-4xl md:text-5xl font-bold mb-6"
        >
          <span className="text-text">Complete CLI reference</span>
        </motion.h2>
        <motion.p 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
          className="text-xl text-text-muted max-w-3xl mx-auto"
        >
          Every command and flag you need to master keyboard shortcuts
        </motion.p>
      </div>

      <motion.div
        initial={{ opacity: 0, y: 20 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true }}
        className="mb-16"
      >
        <CodeBlock code={codeExample} />
      </motion.div>

      <div className="grid lg:grid-cols-2 gap-12">
        <motion.div
          initial={{ opacity: 0, x: -20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
        >
          <Card className="surface border-border h-full">
            <CardHeader>
              <CardTitle className="flex items-center gap-3 text-2xl">
                <Terminal className="w-6 h-6 text-primary" />
                Primary commands
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              {commands.map((item, index) => (
                <div key={index} className="border-b border-border last:border-b-0 pb-4 last:pb-0">
                  <code className="bg-code-bg px-3 py-2 rounded font-mono text-sm border border-border block mb-2">
                    {item.command}
                  </code>
                  <p className="text-text-muted text-sm">
                    {item.description}
                  </p>
                </div>
              ))}
            </CardContent>
          </Card>
        </motion.div>

        <motion.div
          initial={{ opacity: 0, x: 20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
        >
          <Card className="surface border-border h-full">
            <CardHeader>
              <CardTitle className="flex items-center gap-3 text-2xl">
                <Flag className="w-6 h-6 text-secondary" />
                Core flags
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              {flags.map((item, index) => (
                <div key={index} className="border-b border-border last:border-b-0 pb-4 last:pb-0">
                  <code className="bg-code-bg px-3 py-2 rounded font-mono text-sm border border-border block mb-2">
                    {item.flag}
                  </code>
                  <p className="text-text-muted text-sm">
                    {item.description}
                  </p>
                </div>
              ))}
              <div className="bg-accent/10 border border-accent/20 rounded-lg p-4 mt-6">
                <p className="text-text-muted text-sm">
                  The <code className="bg-code-bg px-1 py-0.5 rounded text-xs">karayaml edit</code> command uses the <code className="bg-code-bg px-1 py-0.5 rounded text-xs">$EDITOR</code> environment variable or defaults to VS Code to open the YAML file.
                </p>
              </div>
            </CardContent>
          </Card>
        </motion.div>
      </div>
    </section>
  );
}
