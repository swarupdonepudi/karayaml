import React from "react";
import { motion } from "framer-motion";
import { Badge } from "@/components/ui/badge";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { AlertCircle, Key } from "lucide-react";
import CodeBlock from "./CodeBlock";

export default function KeyUsageInfo() {
  const gotchaCode = `- key: "y"
  file: /System/Applications/Notes.app
- key: "n"
  file: /System/Applications/Numbers.app`;

  return (
    <section id="key-usage" className="max-w-7xl mx-auto px-6 py-32">
      <div className="text-center mb-20">
        <motion.h2 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-4xl md:text-5xl font-bold mb-6"
        >
          Supported Keys &{" "}
          <span className="bg-gradient-to-r from-accent to-orange-500 bg-clip-text text-transparent">
            Gotchas
          </span>
        </motion.h2>
        <motion.p 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
          className="text-xl text-text-muted max-w-3xl mx-auto"
        >
          A quick reference for valid shortcut keys and common configuration pitfalls.
        </motion.p>
      </div>

      <div className="grid lg:grid-cols-2 gap-12 max-w-6xl mx-auto">
        <motion.div
          initial={{ opacity: 0, x: -20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
        >
          <h3 className="text-2xl font-bold mb-6 flex items-center gap-3">
            <Key className="w-6 h-6 text-primary" />
            Valid Key Mappings
          </h3>
          <p className="text-text-muted mb-6">You can use any of the following keys for your shortcuts:</p>
          <div className="space-y-4">
            <div>
              <h4 className="font-semibold mb-2">Alphanumeric Keys</h4>
              <div className="flex flex-wrap gap-2">
                {'abcdefghijklmnopqrstuvwxyz'.split('').map(char => <Badge key={char} variant="secondary" className="font-mono text-lg">{char}</Badge>)}
                {'0123456789'.split('').map(char => <Badge key={char} variant="secondary" className="font-mono text-lg">{char}</Badge>)}
              </div>
            </div>
            <div>
              <h4 className="font-semibold mb-2">Function Keys</h4>
              <div className="flex flex-wrap gap-2">
                {Array.from({ length: 12 }, (_, i) => `F${i + 1}`).map(fkey => <Badge key={fkey} variant="secondary" className="font-mono text-lg">{fkey}</Badge>)}
              </div>
            </div>
          </div>
        </motion.div>

        <motion.div
          initial={{ opacity: 0, x: 20 }}
          whileInView={{ opacity: 1, x: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
        >
          <h3 className="text-2xl font-bold mb-6 flex items-center gap-3">
            <AlertCircle className="w-6 h-6 text-accent" />
            Important YAML Gotcha
          </h3>
          <Alert className="bg-accent/10 border-accent/20">
            <AlertTitle className="font-bold">Quoting 'y' and 'n' keys</AlertTitle>
            <AlertDescription className="text-text-muted leading-relaxed mt-2">
              In YAML, `y`, `yes`, `n`, and `no` are automatically interpreted as boolean `true` or `false`. To use the letters 'y' or 'n' as shortcut keys, you <strong className="text-text">must</strong> wrap them in double quotes.
            </AlertDescription>
          </Alert>
          <CodeBlock
            code={gotchaCode}
            language="yaml"
            filename="Correct Usage"
            className="mt-6"
          />
        </motion.div>
      </div>
    </section>
  );
}