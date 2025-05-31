# 🎯 SRE Key Concepts and Best Practices

---

### 🕹️ SLO (Service Level Objective)
- Explicit target reliability expressed as a percentage.  
- Should be **realistic, achievable**, and **reviewed regularly**.

---

### 📏 SLI (Service Level Indicator)
- Examples include **request latency**, **error rate**, and other measurable metrics.

---

### ⚠️ Error Budget
- The allowable threshold of unreliability within your SLO.  
- Guides how much failure or downtime is acceptable before action is needed.

---

### 🚨 Alerting Based on SLO
- Alerts should trigger on **SLO violations** (not just low-level system errors).  
- Overall SLO = API Gateway SLO × Lambda SLO × DB SLO (multiplicative impact).

---

### 📊 Monitoring

**Whitebox Monitoring**  
- Internal system metrics such as logs, JVM stats, and other in-depth insights.

**Blackbox Monitoring**  
- Observes externally visible behavior, simulating user experience.

**Best Practices:**  
- Focus on rules that reliably catch real incidents.  
- Remove rarely used or non-actionable signals.  
- Avoid combining unrelated functions (like profiling) with monitoring to reduce complexity.  
- Keep systems simple, loosely coupled, and easy to understand.

---

### 🔢 Essential Metrics to Track

1. Invocation count  
2. Error rate and error count  
3. Duration/Latency  
4. Memory usage  
5. Cold start frequency  
6. API Gateway metrics  

---

### 🚨 Alerting Guidelines

- Tie alerts to **user-impacting symptoms**, not just technical errors.  
- Aim to **only wake humans** for urgent, meaningful actions.  
- Focus on user experience metrics like **latency, error rates, and availability**.

---

### 🤖 Elimination of Toil

1. Identify and measure toil.  
2. Automate repetitive tasks.  
3. Improve architecture to reduce toil sources.  
4. Delegate tasks to on-call teams as appropriate.

---

### 🎯 Simplicity

- Inculcate **simplicity end-to-end** in system design and operations.

---

### 🔄 Why Use Decoupling with Queues?

- **Decoupling:** Producers and consumers are independent and scale separately.  
- **Resilience:** Queues buffer requests if functions are busy or down.  
- **Scalability:** Functions scale based on queue depth and demand.  
- **Flexibility:** Supports triggering functions from a wide variety of event sources.

---

### 🚨 Incident Response Best Practices

1. Clear role and command structure.  
2. Early and explicit incident declaration.  
3. Centralized communication channels.  
4. Live documentation during incident handling.

**Additional Tips:**  
- Write postmortems for all significant incidents and share widely.  
- Use collaborative tools and live docs during reviews.  
- Reward effective postmortem authorship and incident handling.  
- Regularly refine the postmortem process with team feedback.  
- Practice training exercises like the **"Wheel of Misfortune"** to improve incident response.

---

### ⚙️ Non-Abstract Large System Design (NALSD)  
An SRE methodology focused on practical, real-world system design emphasizing:

- **Practicality over abstraction**  
- Iterative, requirement-driven processes  
- Reliability as a first-class feature  
- Capacity planning and resource estimation  
- Component isolation and fault tolerance  
- Early SRE involvement in design phases

---

### 🔧 Configuration Design Principles

1. Treat configuration as code.  
2. Separate concerns clearly.  
3. Validate and test configurations.  
4. Perform gradual rollouts.  
5. Centralize and declare management.  
6. Limit scope and complexity.  
7. Implement access controls and auditing.  
8. Automate toil away wherever possible.

---

### 🐤 Canary Releases

- A deployment strategy where new changes are rolled out to a **small subset of production users (the “canary”)** before a full rollout.  
- Helps detect issues early with minimal impact on the overall system.
