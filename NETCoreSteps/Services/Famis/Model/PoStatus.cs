using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace Famis.Model
{
    public class PoStatus
    {
        public string Id { get; set; }
        public string Name { get; set; }
        public bool Active { get; set; }
        public string TabOrder { get; set; }
        public bool DefaultOpenStatusFlag { get; set; }
        public bool EnteredInErrorFlag { get; set; }
        public bool DefaultApprovedStatusFlag { get; set; }
        public bool DefaultDeclinedStatusFlag { get; set; }
        public bool CanceledStatusFlag { get; set; }
        public bool DefaultApprovalWaitingStatusFlag { get; set; }
        public bool DefaultClosedStatusFlag { get; set; }
        public bool DefaultAlterationDeclinedStatusFlag { get; set; }
        public bool AlterationInProgressStatusFlag { get; set; }
        public bool DefaultAutoUpdatePartialFlag { get; set; }
        public bool DefaultAutoUpdateFullFlag { get; set; }
        public bool LockedStatusFlag { get; set; }
    }
}