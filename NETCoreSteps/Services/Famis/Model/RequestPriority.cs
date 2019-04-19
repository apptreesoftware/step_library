using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model {
    public class RequestPriority {
        public bool Active { get; set; }
        public int? DefaultSlaCompletionTime { get; set; }
        public int? DefaultSlaResponseTime { get; set; }
        public string ExternalId { get; set; }
        public int? Id { get; set; }
        public string Name { get; set; }
        public bool OverrideSlaServiceHours { get; set; }
        public bool PushWorkOrdersToExternalSystem { get; set; }
        public int? ResponseLimit { get; set; }
        public string RequestPriorityListId { get; set; }
        public bool ScheduledWorkOrdersOnly { get; set; }
        public int? TabOrder { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedByExternalId { get; set; }
        public int? UpdatedById { get; set; }
        public string Level { get; set; }
        public bool EmergencyEscalation { get; set; }
    }
}
