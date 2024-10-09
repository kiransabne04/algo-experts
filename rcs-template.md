## RCA (Root Cause Analysis) Document Template

When production bugs occur due to the deployment of a developer's branch, an **RCA (Root Cause Analysis) Document** must be created by the developer responsible for the branch. The document provides a detailed account of the issue, the steps taken to resolve it, and recommendations for future prevention.

**Location**: The RCA document should be placed within the relevant microservice directory:
```
/microservice_name/ReadmeFiles/RCA Report/
```
### Example RCA Report

**Task**:  
The customer is unable to log in to the application.
#### **Steps Taken**:

1.  Verified that the customer is using the correct username and password.
2.  Checked the application logs for any errors or anomalies.
3.  Attempted to log in to the application using the customer's credentials.
4.  Contacted the customer to gather more information about the issue.

#### **Findings**:

-   The customer was using the correct username and password.
-   There were no errors recorded in the application logs.
-   I successfully logged into the application with the customer's credentials.
-   The customer was using a VPN to connect to the internet.

#### **Solution**:

-   Instructed the customer to temporarily disable their VPN and attempt to log in again.
-   After disabling the VPN, the customer successfully logged in.

#### **Recommendations**:

-   Advise the customer to reach out to their VPN provider to check for compatibility issues between the VPN and the application.
-   Consider adding a message in the application to notify users that VPN connections could interfere with the login process.

#### **Conclusion**:

The issue was caused by the customer using a VPN, which likely interfered with the login process. After disabling the VPN, the customer was able to log in successfully. We recommend that the customer contacts their VPN provider for further assistance if the issue persists.

----------

**Template for RCA Reports**:

plaintext

Copy code

`#### Task:
Provide a short description of the task or issue.

#### Steps Taken:
1. List all actions taken to investigate and troubleshoot the issue.
2. Include logs, tests, or steps for recreating the issue.
3. Record communication with the customer or stakeholders (if applicable).

#### Findings:
Summarize what was found during the investigation:
- Misconfigurations
- Incorrect user inputs
- Technical failures
- Dependency issues

#### Solution:
Describe the solution that resolved the issue:
- Fixes implemented (code changes, configurations, etc.).
- Steps taken to apply and verify the solution.

#### Recommendations:
Propose preventive measures for future avoidance:
- Documentation updates
- Monitoring enhancements
- Development practices to avoid recurrence

#### Conclusion:
Summarize the root cause of the issue and the overall impact on the application and customer.` 

----------

### Guidelines for RCA Reports:

-   **Clarity**: Ensure the RCA document is clear and easy to understand by both technical and non-technical team members.
-   **Detail**: Provide enough details to ensure that the issue can be referenced in the future and avoided.
-   **Documentation**: Link to any relevant logs, error reports, or files associated with the RCA.

By following this RCA template, the development team can maintain high-quality incident resolution practices while ensuring transparency and accountability for issues that arise during production deployments.
