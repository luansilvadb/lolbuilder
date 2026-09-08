---
name: grilling
description: Stress-test a plan or design through a dependency-aware, one-question-at-a-time interview with recommendations and evidence checks. Use when the user asks to be grilled, challenged, or interviewed before implementation. Do not use for ordinary implementation or general code-review requests.
---

Stress-test the plan thoroughly, not endlessly. The outcome is shared understanding of the material decisions, constraints, evidence, risks, and next step. Stay in planning mode unless the user separately asks for implementation.

Before asking, inspect the plan, repository instructions, relevant code, documentation, and prior tool or agent results that are available. Resolve discoverable facts yourself. If access fails, name the failed lookup and request the smallest missing artifact; do not ask the user to recall a fact known to live in an inaccessible source.

Maintain a lightweight decision state: resolved decisions, assumptions, constraints, open decisions, dependencies, and supporting evidence. Import that state from any prior handoff. Do not reopen a resolved decision unless new evidence conflicts with it; explain the conflict first.

Choose the highest-leverage unresolved decision whose dependencies are already resolved. Each interview turn must contain exactly one interrogative sentence and address one decision. Give the recommended answer before the question, with the main rationale or tradeoff. A useful default shape is:

```text
Recommendation: <preferred answer and why>
Question: <one decision>?
```

After the user answers, update the decision state and continue down the dependency tree. Treat uncertainty about a discoverable fact as a lookup task, not an interview question.

When using a tool or handing work to another agent, carry forward the resolved decisions, constraints, current open decision, and the evidence requested. Reconcile the returned result into the decision state before asking the next question. Never delegate a choice that requires the user's judgment.

Stop when no material open decision remains, when the remaining uncertainty requires unavailable evidence or authority, or when the user asks to stop. Do not invent another question to prolong the interview. Close without a question using a compact handoff with these labels: `Decisions`, `Evidence and assumptions`, `Remaining risks`, and `Next step`.
